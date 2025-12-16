#!/bin/bash

API_URL="http://localhost:3000"

case "$1" in
    propose)
        TITLE="$2"
        DESC="$3"
        curl -s -X POST "$API_URL/governance/proposals" \
             -H "Content-Type: application/json" \
             -d "{\"title\":\"$TITLE\", \"description\":\"$DESC\", \"creator_id\":\"user1\"}" | jq .
        ;;
    vote)
        curl -s -X POST "$API_URL/governance/proposals/$2/votes" \
             -H "Content-Type: application/json" \
             -d "{\"voter_id\":\"user1\", \"approved\":$3}" | jq .
        ;;
    status)
        curl -s "$API_URL/governance/proposals/$2" | jq .
        ;;
    *)
        echo "Uso: $0 [comando]"
        echo ""
        echo "Comandos:"
        echo "  propose \"título\" \"descripción\"  Crear una nueva propuesta"
        echo "  vote [id_propuesta] [true/false] Votar en una propuesta"
        echo "  status [id_propuesta]          Ver estado de una propuesta"
        ;;
esac
