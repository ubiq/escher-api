package dao

import (
  "log"

  . "github.com/ubiq/escher-api/models"
  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

type EscherDAO struct {
  Server    string
  Database  string
}

var db *mgo.Database

const (
  STATUS = "status"
  PROPOSALS = "proposals"
  AIRDROPS = "airdrops"
  VOTES = "votes"
  CLAIMS = "claims"
)

func (e *EscherDAO) Connect() {
  session, err := mgo.Dial(e.Server)
  if err != nil {
    log.Fatal(err)
  }
  db = session.DB(e.Database)
}

func (e *EscherDAO) Status() (Status, error) {
  var status Status
  err := db.C(STATUS).Find(bson.M{}).One(&status)
  return status, err
}

func (e *EscherDAO) Proposals() ([]Proposal, error) {
  var proposals []Proposal
  err := db.C(PROPOSALS).Find(bson.M{}).All(&proposals)
  return proposals, err
}

func (e *EscherDAO) Airdrops() ([]Airdrop, error) {
  var airdrops []Airdrop
  err := db.C(AIRDROPS).Find(bson.M{}).All(&airdrops)
  return airdrops, err
}

func (e *EscherDAO) Airdrop(contract string) (Airdrop, error) {
  var airdrop Airdrop
  err := db.C(AIRDROPS).Find(bson.M{"contract": contract}).One(&airdrop)
  return airdrop, err
}

func (e *EscherDAO) Proposal(contract string) (Proposal, error) {
  var proposal Proposal
  err := db.C(PROPOSALS).Find(bson.M{"contract": contract}).One(&proposal)
  return proposal, err
}

func (e *EscherDAO) Vote(contract string, address string) (Vote, error) {
  var vote Vote
  err := db.C(VOTES).Find(bson.M{"contract": contract, "address": address}).One(&vote)
  return vote, err
}

func (e *EscherDAO) Votes(contract string) ([]Vote, error) {
  var votes []Vote
  err := db.C(VOTES).Find(bson.M{"contract": contract}).All(&votes)
  return votes, err
}

func (e *EscherDAO) Claim(contract string, address string) (Claim, error) {
  var claim Claim
  err := db.C(CLAIMS).Find(bson.M{"contract": contract, "address": address}).One(&claim)
  return claim, err
}

func (e *EscherDAO) Claims(contract string) ([]Claim, error) {
  var claims []Claim
  err := db.C(CLAIMS).Find(bson.M{"contract": contract}).All(&claims)
  return claims, err
}
