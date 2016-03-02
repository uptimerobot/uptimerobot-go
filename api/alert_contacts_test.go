package api

import "testing"

var testAlertID int

func TestNewAlertContacts(t *testing.T) {
	c := makeClient(t)

	alertContacts := c.AlertContacts()

	var request = NewAlertContactRequest{
		AlertContactType:         AlertContactEmail,
		AlertContactValue:        "testcontact@test.com",
		AlertContactFriendlyName: "testcontact",
	}
	response, err := alertContacts.New(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("No alert contacts response: %v", response)
	}
	t.Logf("Alert Contact ID: %d", response.ID)
	testAlertID = response.ID
}

func TestGetAlertContacts(t *testing.T) {

	if testAlertID == 0 {
		t.Skip("TestDeleteAlertContact has no alert id set")
	}

	c := makeClient(t)

	alertContacts := c.AlertContacts()

	alertContactsResp, err := alertContacts.Get()
	if err != nil {
		t.Fatal(err)
	}

	found := false
	for _, ac := range alertContactsResp.AlertContacts {
		if ac.ID == testAlertID &&
			ac.FriendlyName == "testcontact" &&
			ac.Type == AlertContactEmail &&
			ac.Value == "testcontact@test.com" {

			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Couldnt find created alert contact")
	}

}

func TestDeleteAlertContact(t *testing.T) {

	if testAlertID == 0 {
		t.Skip("TestDeleteAlertContact has no alert id set")
	}

	c := makeClient(t)

	alertContacts := c.AlertContacts()

	var request = DeleteAlertContactRequest{
		ID: testAlertID,
	}
	response, err := alertContacts.Delete(request)
	if err != nil {
		t.Fatal(err)
	}
	if response == nil {
		t.Fatal("No alert contact response: %v", response)
	}
	t.Logf("Alert Contact ID: %d", response.ID)

	if response.ID != testAlertID {
		t.Fatalf("Wrong alert id expected %d but got %d\n", testAlertID, response.ID)
	}
}
