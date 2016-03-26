package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

const PHANTOMJS_IMAGE_NAME string = "wernight/phantomjs"
const READER_BUFFER_SIZE int = 4096

// func main() {
// 	//cmd := exec.Command("/bin/bash", "-c", "docker run --rm -it -v $(pwd)/scrapes/www_1015_com.js:/www_1015_com.js wernight/phantomjs /usr/local/bin/phantomjs /www_1015_com.js")
//
// 	cmd := exec.Command("/bin/bash", "-c", "docker run --rm ubuntu /bin/bash -c 'echo foo'")
//
// 	//cmd := exec.Command("docker", "run", "--rm", "ubuntu", "/bin/bash", "-c", "'echo foo'")
// 	//cmd := exec.Command(""docker run --rm ubuntu /bin/bash -c 'echo foo'")
//
// 	out, err := cmd.Output()
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	//for i, char := range out {
// 	//	fmt.Println(i, string(char))
// 	//}
// 	fmt.Println(string(out))
// }

type Scraper struct {
	Id string
}

func (s *Scraper) Scrape(path string) (stdout []byte, err error) {
	fmt.Printf("eyyyy")
	//args := []string{"run", "-d", "-it", PHANTOMJS_IMAGE_NAME}
	//phantomjs_cmd := fmt.Sprint("/usr/local/bin/phantomjs %s", path)
	phantomjs_cmd := fmt.Sprint("/bin/bash echo foo", path)
	cmd := exec.Command("docker", "exec", s.Id, phantomjs_cmd)

	stdout_reader, stdout_reader_err := cmd.StdoutPipe()
	stderr_reader, stderr_reader_err := cmd.StderrPipe()

	fmt.Printf("1")
	if stdout_reader_err != nil {
		panic(stdout_reader_err)
	}

	fmt.Printf("2")
	if stderr_reader_err != nil {
		panic(stderr_reader_err)
	}

	fmt.Printf("3")
	cmd.Start()

	stdout, stdout_err := ioutil.ReadAll(stdout_reader)
	stderr, stderr_err := ioutil.ReadAll(stderr_reader)

	cmd.Wait()

	fmt.Printf("4")
	if stdout_err != nil {
		panic(stdout_err)
	}
	fmt.Printf("5")

	if stderr_err != nil {
		panic(stderr_err)
	}

	if len(stderr) != 0 {
		err = fmt.Errorf("Scraper Scrape error:\n%s", stderr)
	}

	return stdout, err
}

func (s *Scraper) Stop() (err error) {
	//args := []string{"run", "-d", "-it", PHANTOMJS_IMAGE_NAME}
	cmd := exec.Command("docker", "stop", s.Id)

	stderr_reader, stderr_reader_err := cmd.StderrPipe()

	if stderr_reader_err != nil {
		panic(stderr_reader_err)
	}

	cmd.Start()

	stderr, stderr_err := ioutil.ReadAll(stderr_reader)

	cmd.Wait()

	if stderr_err != nil {
		panic(stderr_err)
	}

	if len(stderr) != 0 {
		err = fmt.Errorf("Scraper Stop error:\n%s", stderr)
	}

	return err
}

func Start() (s *Scraper, err error) {
	//args := []string{"run", "-d", "-it", PHANTOMJS_IMAGE_NAME}
	cmd := exec.Command("docker", "run", "-d", "-it", PHANTOMJS_IMAGE_NAME)

	stdout_reader, stdout_reader_err := cmd.StdoutPipe()
	stderr_reader, stderr_reader_err := cmd.StderrPipe()

	if stdout_reader_err != nil {
		panic(stdout_reader_err)
	}

	if stderr_reader_err != nil {
		panic(stderr_reader_err)
	}

	cmd.Start()

	stdout, stdout_err := ioutil.ReadAll(stdout_reader)
	stderr, stderr_err := ioutil.ReadAll(stderr_reader)

	cmd.Wait()

	if stdout_err != nil {
		panic(stdout_err)
	}

	if stderr_err != nil {
		panic(stderr_err)
	}

	if len(stderr) != 0 {
		//error_str := fmt.Sprintf("Scraper Start error:\n%s", stderr)
		err = fmt.Errorf("Scraper Start error:\n%s", stderr)
		return
	}

	return &Scraper{string(stdout)}, err
}

func main() {
	scraper, scraper_err := Start()

	if scraper_err != nil {
		panic(scraper_err)
	}

	fmt.Println(scraper)

	//exec, exec_err := scraper.Scrape("http://google.com")
	exec, exec_err := scraper.Scrape("http://google.com")

	if exec_err != nil {
	}

	//if exec_err != nil {
	//	panic(exec_err)
	//}

	fmt.Println(exec)
}

// func reader_channel(r io.Reader) {
// 	b := make([]byte{}, READER_BUFFER_SIZE)
// 	c := make(chan byte, READER_BUFFER_SIZE)
//
// 	for s := 0; ; {
//
// 		bytes_read, err := r.Read(b)
// 		size += bytes_read
//
// 		if n == 0 {
// 			close(c)
// 		}
// 	}
//
// }
//
// func scrape(s string) <-chan byte {
// 	// command string
// 	cmd_s := fmt.Sprintf("docker run --rm wernight/phantomjs /bin/bash -c 'echo foo'")
//
// 	// run the command
// 	cmd := exec.Command("/bin/bash", "-c", cmd_s)
//
// 	// create the command channel
// 	cmd_c := make(chan byte)
//
// 	// the reader
// 	reader, reader_err := cmd.StdoutPipe()
//
// 	//go func() {
// 	//	for b = make([]byte{}, READER_BUFFER_SIZE)
// 	//}
// }
