package Domain

import driver "github.com/johnnadratowski/golang-neo4j-bolt-driver"

type Ecosystem struct {
	EcosystemID string
	title string
	description string 
	status string
	date_created string
	update_time string
	privacy string
	number_of_members int
}
type EcoSystemRepository interface {
	createEcosystem(EcoSystemID, title, description, status string) error
	updateEcosystem(EcoSystemID, title, description, status string) error
	destroyEcosystem(EcosystemID string) error
	ApplyToEcoSystem(UserID, EcosystemID string) error
	CancelApplication(UserID, EcosystemID string) error
	acceptUser(UserID, EcosystemID, relationship string) error
	declineUser(UserID, EcosystemID string) error
	removeUser(UserID, EcosystemID string) error
}
func (e *Ecosystem) createEcosystem(EcoSystemID, title, description, status, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`CREATE(
		n:ECOSYSTEM{EcoSystemID:{EcoSystemID}, title:{title}, description:{description}, 
		status:{status}})`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"title": title,
			"description": description,
			"status": status,
		})
		if err != nil {
			return err
		}
	return nil
}
func (e *Ecosystem) destroyEcosystem(EcoSystemID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH(n:ECOSYSTEM{EcoSystemID:{EcoSystemID}}) 
							DETACH DELETE n"`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
		})
		if err != nil {
			return err
		}
	return nil
}
func (e *Ecosystem) addStudent(EcoSystemID, UserID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH (e:ECOSYSTEM),(u:USER)
							WHERE e.EcoSystemID = '{EcoSystemID}' AND u.UserID = '{UserID}'
							CREATE (u)-[r:STUDENT_OF]->(e)"`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"UserID": UserID,
		})
		if err != nil {
			return err
		}
	return nil
}
func (e *Ecosystem) removeStudent(EcoSystemID, UserID, neo4jURL string) error{
	db, err := driver.NewDriver().OpenNeo(neo4jURL)
	defer db.Close()
	_, err = db.ExecNeo(`"MATCH (u:USER{UserID:$UserID})-[r.STUDENT_OF]->(e:ECOSYSTEM{EcoSystemID=$EcoSystemID})
							DELETE r`,
		map[string]interface{}{
			"EcoSystemID": EcoSystemID,
			"UserID": UserID,
		})
		if err != nil {
			return err
		}
	return nil
}
func (e *Ecosystem) ApplyToEcoSystem(userToken, EcosystemID string){}
func (e *Ecosystem) LeaveEcoSystem(userToken, EcosystemID string){}
func (e *Ecosystem) AcceptApplicant(ApplicantToken, ModeratorToken string){}
func (e *Ecosystem) RejectApplicant(ApplicantToken, ModeratorToken string){}
