package Publication

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"reflect"
	"time"
)
type PublicationHandler struct {
	PublicationRepository PublicationRepository
}
type PublicationNeo4jRepository struct {
	Driver neo4j.Driver
}
type Publication struct{
	Title string
	Description string
	publishers []Publisher
	logs []Log
	sponsor []Sponsor
	files []File
	Date_Created string
	status string
	date_created time.Time
	update_time time.Time
}
type Publisher struct{
	first_name string
	last_name string
	userID string
}
type File struct{
	filename string
	filetype string
	filesize string
	data byte
}
type Log struct {
	LogName string
	description string
	author Publisher
	date_created time.Time

}
type Sponsor struct{
	
}
type PublicationRepository interface {

	changeTitle(projectID, title string) error
	changeDescription(projectID, description string) error
	
	setToPrivate(projectID string) error
	setToPublic(projectID string) error

	storePackages(projectID string, files []File) error
	removePackages(projectID string, files []string) error

	createProject(projectID, title, description, date_created, time_updated string) error
	deleteProject(projectID string) error
	
	getPublication(projectID, userID string) error
	listPublications(userID string) error
	
	addPublisher(projectID, userID string) error
	removePublisher(projectID, userID string) error
}

func (p *Publication) GetProjectFields(field, value string) string {
	r := reflect.ValueOf(p)
	f := reflect.Indirect(r).FieldByName(field)
	return f.FieldByName(value).String()
}
func (p *PublicationNeo4jRepository) createProject(userID, projectID, ecosystemID, title, description, date_created, time_updated string) error{
	session := p.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	query := `MATCH (u:USER{UserID:$UserID}) CREATE (u) -[r:CREATED]-> (p:PUBLICATION
		{ProjectID:$ProjectID, Title:$Title, Description:$Description, Date_Created:$Date_Created,
		 Time_Updated:$Time_Updated}) <-[r:SPONSORS]-(e:EcoSystem{EcoSystemID:$EcoSystenID})`
	parameters := map[string]interface{}{
		"UserID": userID,
		"ProjectID": projectID,
		"EcoSystemID": ecosystemID,
		"Title": title,
		"Description": description,
		"Date_Created": date_created,
		"Time_Updated": time_updated,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (p *PublicationNeo4jRepository) deleteProject(projectID string) error{
	session := p.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	query := `MATCH (u:USER{UserID:$UserID}) CREATE (u) -[r:CREATED]-> (p:PUBLICATION
		{ProjectID:$ProjectID, Title:$Title, Description:$Description, Date_Created:$Date_Created,
		 Time_Updated:$Time_Updated})`
	parameters := map[string]interface{}{
		"ProjectID": projectID,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (p *PublicationNeo4jRepository) changeTitle(projectID, title string, time_updated time.Time) error{
	session := p.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	query := `MATCH (p:PUBLICATION{projectID:$ProjectID}) 
				SET p.Title = toString($Title), p.Time_Updated = $Time`
	parameters := map[string]interface{}{
		"ProjectID": projectID,
		"Title": title,
		"Time": time_updated,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (p *PublicationNeo4jRepository) changeDescription(projectID, description string, time_updated time.Time) error{
	session := p.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	query := `MATCH (p:PUBLICATION{projectID:$ProjectID}) 
				SET p.Description = toString($Description), p.Time_Updated = $Time`
	parameters := map[string]interface{}{
		"ProjectID": projectID,
		"Description": description,
		"Time": time_updated,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}