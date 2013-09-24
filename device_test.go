package ssdp

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_ParseDecriptionXml(t *testing.T) {
	descriptionFile := filepath.Join(".", "test_responses", "hue_description.xml")
	fileBytes, err := ioutil.ReadFile(descriptionFile)
	if err != nil {
		t.Fatal("Error reading in stub description.xml.", err)
	}

	deviceDescription, err := decodeDescription(bytes.NewReader(fileBytes))
	if err != nil {
		t.Fatal("Could not decode description.", err)
	}

	assertEqual(t, 1, deviceDescription.SpecVersion.Major, "SpecVersion.Major")
	assertEqual(t, 0, deviceDescription.SpecVersion.Minor, "SpecVersion.Minor")
	assertEqual(t, "http://192.168.0.21:80/", deviceDescription.UrlBase, "UrlBase")
	assertEqual(t, "urn:schemas-upnp-org:device:Basic:1", deviceDescription.DeviceType, "DeviceType")
	assertEqual(t, "Philips hue (192.168.0.21)", deviceDescription.FriendlyName, "FriendlyName")
	assertEqual(t, "Royal Philips Electronics", deviceDescription.Manufacturer, "Manufacturer")
	assertEqual(t, "http://www.philips.com", deviceDescription.ManufacturerUrl, "ManufacturerUrl")
	assertEqual(t, "Philips hue Personal Wireless Lighting", deviceDescription.ModelDescription, "ModelDescription")
	assertEqual(t, "Philips hue bridge 2012", deviceDescription.ModelName, "ModelName")
	assertEqual(t, "1000000000000", deviceDescription.ModelNumber, "ModelNumber")
	assertEqual(t, "http://www.meethue.com", deviceDescription.ModelUrl, "ModelUrl")
	assertEqual(t, "93eadbeef13", deviceDescription.SerialNumber, "SerialNumber")
	assertEqual(t, "uuid:01234567-89ab-cdef-0123-456789abcdef", deviceDescription.Udn, "Udn")
	assertEqual(t, "", deviceDescription.Upc, "Upc")
	assertEqual(t, "index.html", deviceDescription.PresentationUrl, "PresentationUrl")

	icons := deviceDescription.Icons
	assertEqual(t, 2, len(icons), "len(icons)")
	assertEqual(t, "image/png", icons[0].MimeType, "icons.MimeType")
	assertEqual(t, 48, icons[0].Width, "icons.Width")
	assertEqual(t, 48, icons[0].Height, "icons.Height")
	assertEqual(t, 24, icons[0].Depth, "icons.Depth")
	assertEqual(t, "hue_logo_0.png", icons[0].Url, "icons.Url")
}
