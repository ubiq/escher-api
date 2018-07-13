package models

type Block struct {
	Number       uint64  `bson:"number" json:"number"`
	Miner        string  `bson:"miner" json:"miner"`
	Timestamp    uint64  `bson:"timestamp" json:"timestamp"`
	Transactions uint64  `bson:"transactions" json:"transactions"`
	Reward       string  `bson:"reward" json:"reward"`
}

type Status struct {
	Block        *Block  `bson:"block" json:"block"`
	TotalUbiq    string  `bson:"totalUbiq" json:"totalUbiq"`
	TotalEscher  string  `bson:"totalEscher" json:"totalEscher"`
}

type Candidate struct {
	Title        string  `bson:"title" json:"title"`
	Index        uint64  `bson:"index" json:"index"`
	TotalWeight  string  `bson:"totalWeight" json:"totalWeight"`
	TotalVotes   uint64  `bson:"totalVotes" json:"totalVotes"`
}

type Proposal struct {
	Title        string  `bson:"title" json:"title"`
	Type         string  `bson:"type" json:"type"`
	Contract     string  `bson:"contract" json:"contract"`
	StartBlock   uint64  `bson:"startBlock" json:"startBlock"`
	EndBlock     uint64  `bson:"endBlock" json:"endBlock"`
	TotalWeight  string  `bson:"totalWeight" json:"totalWeight"`
	Candidates   []Candidate `bson:"candidates" json:"candidates"`
	Data         string  `bson:"data" json:"data"`
	Uip          string  `bson:"uip" json:"uip"`
}

type Airdrop struct {
	Title        string  `bson:"title" json:"title"`
	Contract     string  `bson:"contract" json:"contract"`
	Multiplier   uint64  `bson:"multiplier" json:"multiplier"`
	StartBlock   uint64  `bson:"startBlock" json:"startBlock"`
	EndBlock     uint64  `bson:"endBlock" json:"endBlock"`
	TotalClaimed string  `bson:"totalClaimed" json:"totalClaimed"`
}

type Claim struct {
	Address      string  `bson:"address" json:"address"`
	Balance      string  `bson:"balance" json:"balance"`
	Contract     string  `bson:"contract" json:"contract"`
}

type Vote struct {
	Address      string  `bson:"address" json:"address"`
	Balance      string  `bson:"balance" json:"balance"`
	Contract     string  `bson:"contract" json:"contract"`
	Candidate    uint64  `bson:"candidate" json:"candidate"`
}
