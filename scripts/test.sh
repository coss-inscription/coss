export ALICE=$(cossd keys show alice --address)
export BOB=$(cossd keys show bob --address)
echo $ALICE
echo $BOB

echo '🚩 create token-admin to bob'
sleep 1
cossd tx vault create-token-admin $BOB --from alice --chain-id=coss --yes   
sleep 1
cossd query vault show-token-admin
sleep 3
echo '🚩 create token ucoss from alice'
sleep 1
cossd tx vault create-token ucoss 'Inscription on Cosmos' COSS 6 'https://coss.ink' 21000000000000000 --from alice  --chain-id=coss --yes
sleep 1
cossd query vault show-token ucoss
sleep 3
echo '🚩 audit token ucoss from bob(token-admin)'
sleep 1
cossd tx vault audit-token ucoss true --from bob --chain-id=coss --yes
sleep 1
cossd query vault show-token ucoss
sleep 3
echo '🚩 mint 1000000ucoss to bob'
sleep 1
cossd tx vault convert-ins-to-token 1000000ucoss $BOB --from alice --chain-id=coss --yes  
sleep 1
cossd query vault show-token ucoss
cossd query bank balances $BOB
sleep 3
echo '🚩 burn 1000000ucoss from bob'
sleep 1
cossd tx vault convert-token-to-ins 1000000ucoss --from bob --chain-id=coss --yes
sleep 1
cossd query bank balances $BOB
cossd query vault show-token ucoss
# cossd tx vault update-token ucoss 'Inscription on Cosmos1' COSS 6 'https://coss.ink1' 21000000000000000 --from alice  --chain-id=coss --yes