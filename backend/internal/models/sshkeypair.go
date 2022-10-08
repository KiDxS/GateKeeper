package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type SSHKeyPair struct {
	ID      int    `json:"id"`
	Label   string `json:"label"`
	PubKey  string `json:"pubKey"`
	PrivKey string `json:"privKey"`
}

func (keypair *SSHKeyPair) QuerySSHKeyPair(id int) error {
	db := connect()
	stm, err := db.Prepare("SELECT * FROM sshpair WHERE id = ?")
	if err != nil {
		return err
	}
	err = stm.QueryRow(id).Scan(&keypair.ID, &keypair.Label, &keypair.PubKey, &keypair.PrivKey)

	if err != nil {
		return err
	}

	return nil
}

func (keypair *SSHKeyPair) InsertSSHPairKey(label, pubKey, privKey string) error {
	db := connect()
	stm, err := db.Prepare("INSERT INTO sshpair (label, pub_key, priv_key) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stm.Exec(label, pubKey, privKey)
	if err != nil {
		return err
	}
	return nil
}

func (keypair *SSHKeyPair) QuerySSHKeyPairLabels() ([]string, error) {
	db := connect()
	labels := []string{}
	rows, err := db.Query("SELECT label from sshpair")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&keypair.Label)
		labels = append(labels, keypair.Label)
	}
	return labels, nil
}

func (keypair *SSHKeyPair) DeleteSSHKeyPair(id int) error {
	db := connect()
	err := keypair.QuerySSHKeyPair(id)
	if err != nil {
		return err
	}
	stm, err := db.Prepare("DELETE FROM sshpair WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stm.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (keypair *SSHKeyPair) UpdateSSHKeyPairLabel(id int, label string) error {
	db := connect()
	err := keypair.QuerySSHKeyPair(id)
	if err != nil {
		return err
	}
	stm, err := db.Prepare("UPDATE sshpair SET label = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stm.Exec(label, id)
	if err != nil {
		return err
	}
	err = keypair.QuerySSHKeyPair(id)
	if err != nil {
		return nil
	}
	return nil
}
