package model

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type View interface {
	Render(out io.Writer, name string, data interface{}) error
}