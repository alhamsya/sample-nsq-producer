package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

//ProjectRequest inside data
type ProjectRequest struct {
	ID               int64     `json:"id"`
	ProjectNumber    string    `json:"project_number"`
	Description      string    `json:"description"`
	ApprovalStatus   int       `json:"approval_status"`
	ProjectStatus    string    `json:"project_status"`
	Name             string    `json:"name"`
	ParentProjectID  int       `json:"parent_project_id,omitempty"`
	SecurityGroup    string    `json:"security_group"`
	CategoryID       int       `json:"category_id"`
	BusinessUser     string    `json:"business_user"`
	ContributionType string    `json:"contribution_type"`
	StartDate        string    `json:"start_date"`
	TargetSaving     float64   `json:"target_saving"`
	SupplierRefID    []string  `json:"supplier_ref_id"`
	CreatedAt        time.Time `json:"created_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedBy        string    `json:"updated_by"`
}

//DataConsumer data send to NSQ
type DataConsumer struct {
	EmailTo        []string
	EmailCC        []string
	ProjectRequest ProjectRequest `json:"project_request"`
}

func main() {
	publisher()
}

func publisher() {
	const (
		address = "127.0.0.1:4150"
		topic   = "ent_co_notification"
	)

	config := nsq.NewConfig()
	p, err := nsq.NewProducer(address, config)
	if err != nil {
		log.Panic("[ERR] fail setup producer NSQ :", err)
	}

	data := &DataConsumer{
		EmailTo: []string{"alhamsya@gmail.com"},
		EmailCC: nil,
		ProjectRequest: ProjectRequest{
			ID:               1,
			ProjectNumber:    "tes 1",
			Description:      "",
			ApprovalStatus:   0,
			ProjectStatus:    "",
			Name:             "tes name",
			ParentProjectID:  0,
			SecurityGroup:    "",
			CategoryID:       0,
			BusinessUser:     "",
			ContributionType: "tes",
			StartDate:        "",
			TargetSaving:     0,
			SupplierRefID:    nil,
			CreatedAt:        time.Time{},
			CreatedBy:        "",
			UpdatedBy:        "",
		},
	}

	dataRequest, err := json.Marshal(data)
	if err != nil {
		log.Fatal("[ERR] fail cannot json marshal :", err)
	}

	fmt.Println(string(dataRequest))

	err = p.Publish(topic, []byte(dataRequest))
	if err != nil {
		log.Panic("[ERR] fail publish data to NSQ :", err)
	}

	p.Stop()
}
