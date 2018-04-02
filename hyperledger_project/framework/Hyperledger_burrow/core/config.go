// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// config.go keeps explicit structures on the runtime configuration of
// burrow and all modules.  It loads these from the Viper configuration
// loaded in `definitions.Do`

package core

import (
	"fmt"
	"os"
	"path"

	"github.com/hyperledger/burrow/config"
	"github.com/hyperledger/burrow/definitions"
	lconfig "github.com/hyperledger/burrow/logging/config"
	"github.com/hyperledger/burrow/server"
	"github.com/hyperledger/burrow/util"
	"github.com/spf13/viper"
)

// LoadConsensusModuleConfig wraps specifically for the consensus module
func LoadConsensusModuleConfig(do *definitions.Do) (*config.ModuleConfig, error) {
	return loadModuleConfigFromDo(do, "consensus")
}

// LoadApplicationManagerModuleConfig wraps specifically for the application
// manager
func LoadApplicationManagerModuleConfig(do *definitions.Do) (*config.ModuleConfig, error) {
	return loadModuleConfigFromDo(do, "manager")
}

func loadModuleConfigFromDo(do *definitions.Do, module string) (*config.ModuleConfig, error) {
	return LoadModuleConfig(do.Config, do.WorkDir, do.DataDir,
		do.GenesisFile, do.ChainId, module)
}

// Generic Module loader for configuration information
func LoadModuleConfig(conf *viper.Viper, rootWorkDir, rootDataDir,
	genesisFile, chainId, module string) (*config.ModuleConfig, error) {
	moduleName := conf.GetString("chain." + module + ".name")
	// set up the directory structure for the module inside the data directory
	workDir := path.Join(rootDataDir, conf.GetString("chain."+module+
		".relative_root"))
	if err := util.EnsureDir(workDir, os.ModePerm); err != nil {
		return nil,
			fmt.Errorf("Failed to create module root directory %s.", workDir)
	}
	dataDir := path.Join(workDir, "data")
	if err := util.EnsureDir(dataDir, os.ModePerm); err != nil {
		return nil,
			fmt.Errorf("Failed to create module data directory %s.", dataDir)
	}
	// load configuration subtree for module
	if !conf.IsSet(moduleName) {
		return nil, fmt.Errorf("Failed to read configuration section for %s",
			moduleName)
	}
	subConfig, err := config.ViperSubConfig(conf, moduleName)
	if subConfig == nil {
		return nil, fmt.Errorf("Failed to read configuration section for %s: %s",
			moduleName, err)
	}

	return &config.ModuleConfig{
		Module:      module,
		Name:        moduleName,
		WorkDir:     workDir,
		DataDir:     dataDir,
		RootDir:     rootWorkDir, // burrow's working directory
		ChainId:     chainId,
		GenesisFile: genesisFile,
		Config:      subConfig,
	}, nil
}

// Load the ServerConfig from commandline Do object
func LoadServerConfigFromDo(do *definitions.Do) (*server.ServerConfig, error) {
	// load configuration subtree for servers
	return LoadServerConfig(do.ChainId, do.Config)
}

// Load the ServerConfig from root Viper config, fixing the ChainId
func LoadServerConfig(chainId string, rootConfig *viper.Viper) (*server.ServerConfig, error) {
	subConfig, err := config.ViperSubConfig(rootConfig, "servers")
	if err != nil {
		return nil, err
	}
	serverConfig, err := server.ReadServerConfig(subConfig)
	if err != nil {
		return nil, err
	}
	serverConfig.ChainId = chainId
	return serverConfig, err
}

func LoadLoggingConfigFromDo(do *definitions.Do) (*lconfig.LoggingConfig, error) {
	if !do.Config.IsSet("logging") {
		return nil, nil
	}
	loggingConfigMap := do.Config.GetStringMap("logging")
	return lconfig.LoggingConfigFromMap(loggingConfigMap)
}

func LoadLoggingConfigFromClientDo(do *definitions.ClientDo) (*lconfig.LoggingConfig, error) {
	loggingConfig := lconfig.DefaultClientLoggingConfig()
	return loggingConfig, nil
}
