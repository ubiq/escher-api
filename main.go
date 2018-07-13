package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
	"github.com/gorilla/mux"
	. "github.com/ubiq/escher-api/config"
	. "github.com/ubiq/escher-api/dao"
)

var config_ = Config{}
var dao_ = EscherDAO{}
var port string

func getStatus(w http.ResponseWriter, r *http.Request) {
	// returns total escher supply, total ubiq supply, and latest block
	status, err := dao_.Status()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, status)
}

func getProposals(w http.ResponseWriter, r *http.Request) {
	// returns proposals
	proposals, err := dao_.Proposals()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, proposals)
}

func getAirdrops(w http.ResponseWriter, r *http.Request) {
	// airdrops
	airdrops, err := dao_.Airdrops()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, airdrops)
}

func getAirdrop(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	airdrop, err := dao_.Airdrop(params["contract"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract Address")
		return
	}
	respondWithJson(w, http.StatusOK, airdrop)
}

func getProposal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	proposal, err := dao_.Proposal(params["contract"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract Address")
		return
	}
	respondWithJson(w, http.StatusOK, proposal)
}

func getVote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	vote, err := dao_.Vote(params["contract"], params["address"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract Or Address")
		return
	}
	respondWithJson(w, http.StatusOK, vote)
}

func getClaim(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	claim, err := dao_.Claim(params["contract"], params["address"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Contract Or Address")
		return
	}
	respondWithJson(w, http.StatusOK, claim)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config_.Read()

  port = config_.Port
	dao_.Server = config_.Server
	dao_.Database = config_.Database
	dao_.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/status", getStatus).Methods("GET")
  r.HandleFunc("/airdrops", getAirdrops).Methods("GET")
  r.HandleFunc("/proposals", getProposals).Methods("GET")
	r.HandleFunc("/airdrop/{contract}", getAirdrop).Methods("GET")
  r.HandleFunc("/proposal/{contract}", getProposal).Methods("GET")
  r.HandleFunc("/claim/{contract}/{address}", getClaim).Methods("GET")
  r.HandleFunc("/vote/{contract}/{address}", getVote).Methods("GET")

  handler := cors.Default().Handler(r)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
