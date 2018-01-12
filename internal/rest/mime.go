package rest

import (
	"io/ioutil"

	"github.com/emicklei/go-restful"
	"gopkg.in/yaml.v2"
)

type YamlReaderWriter struct {
	contentType string
}

func NewYamlReaderWriter(contentType string) restful.EntityReaderWriter {
	return YamlReaderWriter{contentType: contentType}
}

// Read a serialized version of the value from the request.
// The Request may have a decompressing reader. Depends on Content-Encoding.
func (e YamlReaderWriter) Read(req *restful.Request, v interface{}) error {
	defer req.Request.Body.Close()
	bytes, err := ioutil.ReadAll(req.Request.Body)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, v)
	return err
}

// Write a serialized version of the value on the response.
// The Response may have a compressing writer. Depends on Accept-Encoding.
// status should be a valid Http Status code
func (e YamlReaderWriter) Write(resp *restful.Response, status int, v interface{}) error {
	bytes, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	resp.WriteHeader(status)
	_, err = resp.Write(bytes)
	return err
}
