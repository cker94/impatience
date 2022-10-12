package impatience

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
)

struct savedStock {
  List []string    `json:"list"`
  Pos uint8        `json:"pos"`
  LoopCount uint8  `json:"loopCount"`
}

struct savedGame {
  Stock savedStock       `json:"stock"`
  Tableau [7][]string    `json:"tableau"`
  Foundations[4][]string `json:"foundations"`
}

// TODO: Load game from file
func LoadGame(filepath string) *Game, error {
  var save savedGame
  // Open file.
  if file, err := os.Open(filepath); err != nil {
    return nil, err
  } else {
    // Whatever happens, close file when function exits.
    defer file.Close()
    // Read file into memory.
    if contents, err := io.ReadAll(file); err != nil {
      return nil, err
    } else {
      // TODO: Add TOML support
      // Parse JSON into Go struct.
      json.Unmarshal(contents, &save)
    }
  }
  
  game := NewGame()
  // TODO: convert save data to actual game state.
  
  return game, nil
}
