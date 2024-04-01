package agent_test

import (
	"bytes"
	"net/http"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/daytonaio/daytona/pkg/agent"
	"github.com/daytonaio/daytona/pkg/agent/config"
	"github.com/daytonaio/daytona/pkg/serverapiclient"
	"github.com/daytonaio/daytona/pkg/types"
	"github.com/gin-gonic/gin"
)

var mockConfig = &config.Config{
	WorkspaceId: "workspace-test",
	ProjectDir:  "/project/helloworld",
	ProjectName: "helloworld",
	Server: config.DaytonaServerConfig{
		Url:    "http://localhost:3000",
		ApiKey: "test-api-key",
	},
}

type MockGitService struct {
	returnRepoExists bool
}

func (m *MockGitService) CloneRepository(project *serverapiclient.Project, authToken *string) error {
	// Implement the mock behavior here
	return nil
}

func (m *MockGitService) RepositoryExists(project *serverapiclient.Project) (bool, error) {
	return m.returnRepoExists, nil
}

func (m *MockGitService) SetGitConfig(userData *serverapiclient.GitUserData) error {
	// Implement the mock behavior here
	return nil
}

type MockRestServer *http.Server

func NewMockRestServer() *http.Server {
	router := gin.Default()
	serverController := router.Group("/server")
	{
		serverController.GET("/config", func(ctx *gin.Context) {
			ctx.JSON(200, &types.ServerConfig{
				ProvidersDir:      "",
				RegistryUrl:       "",
				GitProviders:      []types.GitProvider{},
				Id:                "",
				ServerDownloadUrl: "",
				ApiPort:           3000,
				TailscalePort:     4000,
				TargetsFilePath:   "",
				BinariesPath:      "",
			})
		})
		serverController.POST("/network-key", func(ctx *gin.Context) {
			ctx.JSON(200, &types.NetworkKey{Key: "test-key"})
		})
	}

	workspaceController := router.Group("/workspace")
	{
		workspaceController.GET("/:workspaceId", func(ctx *gin.Context) {
			workspaceId := "workspace-test"
			workspaceName := "test-workspace"
			projectName := "helloworld"
			projectRepoUrl := "https://github.com/go-training/helloworld"
			wrksp := serverapiclient.WorkspaceDTO{
				Id:   &workspaceId,
				Name: &workspaceName,
				Projects: []serverapiclient.Project{
					{
						Name: &projectName,
						Repository: &serverapiclient.Repository{
							Url: &projectRepoUrl,
						},
					},
				},
			}
			ctx.JSON(http.StatusOK, wrksp)
		})
	}

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	return server
}

type MockSshServer struct{}

func (m *MockSshServer) Start() error {
	// Implement the mock behavior here
	return nil
}

type MockTailscaleServer struct{}

func (m *MockTailscaleServer) Start() error {
	// Implement the mock behavior here
	return nil
}

func TestStart(t *testing.T) {
	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	apiServer := NewMockRestServer()
	defer apiServer.Close()
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Create a new Agent instance
	a := &agent.Agent{
		Config: mockConfig,
		Git: &MockGitService{
			returnRepoExists: false,
		},
		Ssh:       &MockSshServer{},
		Tailscale: &MockTailscaleServer{},
	}

	// Call the Start method
	err := a.Start()

	// Check if the error is not nil
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestSkipCloneIfRepoExists(t *testing.T) {
	buf := bytes.Buffer{}
	log.SetOutput(&buf)

	apiServer := NewMockRestServer()
	defer apiServer.Close()
	go func() {
		if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Create a new Agent instance
	a := &agent.Agent{
		Config: mockConfig,
		Git: &MockGitService{
			returnRepoExists: true,
		},
		Ssh:       &MockSshServer{},
		Tailscale: &MockTailscaleServer{},
	}

	// Call the Start method
	err := a.Start()

	// Check if the error is not nil
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !bytes.Contains(buf.Bytes(), []byte("Repository already exists. Skipping clone...")) {
		t.Errorf("Expected log to contain 'Repository already exists. Skipping clone...', got '%s'", buf.String())
	}
}
