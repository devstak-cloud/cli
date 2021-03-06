package pluginaction_test

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	. "code.cloudfoundry.org/cli/actor/pluginaction"
	"code.cloudfoundry.org/cli/actor/pluginaction/pluginactionfakes"
	"code.cloudfoundry.org/cli/util/configv3"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UninstallPlugin", func() {
	var (
		actor                 *Actor
		binaryPath            string
		fakeConfig            *pluginactionfakes.FakeConfig
		fakePluginUninstaller *pluginactionfakes.FakePluginUninstaller
		pluginHome            string
		err                   error
	)

	BeforeEach(func() {
		pluginHome, err = ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())

		binaryPath = filepath.Join(pluginHome, "banana-faceman")
		err = ioutil.WriteFile(binaryPath, nil, 0600)
		Expect(err).ToNot(HaveOccurred())

		plugins := map[string]configv3.Plugin{
			"some-plugin": {
				Location: binaryPath,
			},
		}

		fakePluginUninstaller = new(pluginactionfakes.FakePluginUninstaller)
		fakeConfig = new(pluginactionfakes.FakeConfig)
		actor = NewActor(fakeConfig)

		fakeConfig.PluginsReturns(plugins)
	})

	AfterEach(func() {
		os.RemoveAll(pluginHome)
	})

	Context("when the plugin does not exist", func() {
		It("returns a PluginNotFoundError", func() {
			err := actor.UninstallPlugin(fakePluginUninstaller, "some-non-existent-plugin")
			Expect(err).To(MatchError(PluginNotFoundError{Name: "some-non-existent-plugin"}))
		})
	})

	Context("when the plugin exists", func() {
		Context("when no errors are encountered", func() {
			It("runs the plugin cleanup, deletes the binary and removes the plugin config", func() {
				err := actor.UninstallPlugin(fakePluginUninstaller, "some-plugin")
				Expect(err).ToNot(HaveOccurred())

				Expect(fakeConfig.PluginsCallCount()).To(Equal(1))

				Expect(fakePluginUninstaller.UninstallCallCount()).To(Equal(1))
				Expect(fakePluginUninstaller.UninstallArgsForCall(0)).To(Equal(binaryPath))

				_, err = os.Stat(binaryPath)
				Expect(os.IsNotExist(err)).To(BeTrue())

				Expect(fakeConfig.RemovePluginCallCount()).To(Equal(1))
				Expect(fakeConfig.RemovePluginArgsForCall(0)).To(Equal("some-plugin"))

				Expect(fakeConfig.WritePluginConfigCallCount()).To(Equal(1))
			})
		})

		Context("when the plugin uninstaller returns an error", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some error")
				fakePluginUninstaller.UninstallReturns(expectedErr)
			})

			It("returns the error and does not delete the binary or remove the plugin config", func() {
				err := actor.UninstallPlugin(fakePluginUninstaller, "some-plugin")
				Expect(err).To(MatchError(expectedErr))

				_, err = os.Stat(binaryPath)
				Expect(err).ToNot(HaveOccurred())

				Expect(fakeConfig.RemovePluginCallCount()).To(Equal(0))
			})
		})

		Context("when deleting the plugin binary returns a 'file does not exist' error", func() {
			BeforeEach(func() {
				err = os.Remove(binaryPath)
				Expect(err).ToNot(HaveOccurred())
			})

			It("does not return the error and removes the plugin config", func() {
				err := actor.UninstallPlugin(fakePluginUninstaller, "some-plugin")
				Expect(err).ToNot(HaveOccurred())

				Expect(fakeConfig.RemovePluginCallCount()).To(Equal(1))
			})
		})

		Context("when deleting the plugin binary returns any other error", func() {
			BeforeEach(func() {
				err = os.Remove(binaryPath)
				Expect(err).ToNot(HaveOccurred())
				err = os.Mkdir(binaryPath, 0700)
				Expect(err).ToNot(HaveOccurred())
				err = ioutil.WriteFile(filepath.Join(binaryPath, "foooooo"), nil, 0500)
				Expect(err).ToNot(HaveOccurred())
			})

			It("does not return the error and removes the plugin config", func() {
				err := actor.UninstallPlugin(fakePluginUninstaller, "some-plugin")
				_, ok := err.(*os.PathError)
				Expect(ok).To(BeTrue())

				Expect(fakeConfig.RemovePluginCallCount()).To(Equal(0))
				Expect(fakeConfig.WritePluginConfigCallCount()).To(Equal(0))
			})
		})

		Context("when writing the config returns an error", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("some plugin config write error")

				fakeConfig.WritePluginConfigReturns(expectedErr)
			})

			It("returns the error", func() {
				err := actor.UninstallPlugin(fakePluginUninstaller, "some-plugin")
				Expect(err).To(MatchError(expectedErr))
			})
		})
	})
})
