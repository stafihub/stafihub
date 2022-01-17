package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func parseMetadataFlags(fs *pflag.FlagSet) (*banktypes.Metadata, error) {
	md := &banktypes.Metadata{}
	mdFile, _ := fs.GetString(FlagMetadata)

	if mdFile == "" {
		return nil, fmt.Errorf("metadataFile not give")
	}

	contents, err := os.ReadFile(mdFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, md)
	if err != nil {
		return nil, err
	}

	return md, nil
}
