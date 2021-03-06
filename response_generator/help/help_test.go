package help_test

import (
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/votinginfoproject/sms-worker/civic_api"
	"github.com/votinginfoproject/sms-worker/civic_api/fixtures"
	"github.com/votinginfoproject/sms-worker/fake_storage"
	"github.com/votinginfoproject/sms-worker/response_generator"
	"github.com/votinginfoproject/sms-worker/test_helpers"
	"github.com/votinginfoproject/sms-worker/users"
)

func setup() {
	log.SetOutput(ioutil.Discard)
}

var content = testHelpers.GetContent()

func TestHelpWithCommandNotFirstContact(t *testing.T) {
	setup()
	s := fakeStorage.New()
	u := users.New(s)

	time := time.Now().Unix()
	timeString := strconv.FormatInt(time, 10)
	s.CreateItem("+15551235555", map[string]string{"language": "en", "last_contact": timeString})

	c := civicApi.New("", "", "", civicApiFixtures.MakeRequestSuccessFake)
	g := responseGenerator.New(c, u)

  expected := []string{content.Help.Text["en"]["intro"],
                       content.Help.Text["en"]["menu"],
                       content.Help.Text["en"]["languages"]}
	assert.Equal(t, expected, g.Generate("+15551235555", "help", 0))
}

func TestHelpWithCommandFirstContact(t *testing.T) {
	setup()
	s := fakeStorage.New()
	u := users.New(s)

	c := civicApi.New("", "", "", civicApiFixtures.MakeRequestSuccessFake)
	g := responseGenerator.New(c, u)

	expected := []string{content.Intro.Text["en"]["all"]}
	assert.Equal(t, expected, g.Generate("+15551235555", "help", 0))
}
