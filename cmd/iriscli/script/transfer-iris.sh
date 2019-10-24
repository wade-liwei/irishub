cd ~/ibc-testnets

echo "\n*** transfer...\n"
echo 12345678 | iriscli --home ibc-iris/n0/iriscli tx ibcmockbank transfer \
    --src-port bank --src-channel chann-to-gaia \
    --denom uiris --amount 1 \
    --receiver $(gaiacli --home ibc-gaia/n0/gaiacli keys show n0 | jq -r '.address') \
    --source true \
    --from n0 -y -o json \
    --broadcast-mode=block \
    >ibc-iris/n0/result.json
jq -r '.events[1].attributes[5].value' ibc-iris/n0/result.json >ibc-gaia/n0/packet.json
echo "\n*** receive...\n"
sleep 8 && iriscli --home ibc-iris/n0/iriscli q ibc client header -o json >ibc-gaia/n0/header.json
iriscli --home ibc-iris/n0/iriscli q ibc channel packet-proof bank chann-to-gaia \
    $(jq -r '.m_sequence' ibc-gaia/n0/packet.json) \
    $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
    -o json >ibc-gaia/n0/proof.json
echo 12345678 | gaiacli --home ibc-gaia/n0/gaiacli tx ibc client update client-to-iris ibc-gaia/n0/header.json --from n0 -y -o text --broadcast-mode=block
echo 12345678 | gaiacli --home ibc-gaia/n0/gaiacli tx ibcmockbank recv-packet \
    ibc-gaia/n0/packet.json ibc-gaia/n0/proof.json \
    $(jq -r '.value.SignedHeader.header.height' ibc-gaia/n0/header.json) \
    --from n0 -y -o text \
    --broadcast-mode=block
echo "\n*** query account...\n"
gaiacli --home ibc-gaia/n0/gaiacli q account -o text $(gaiacli --home ibc-gaia/n0/gaiacli keys show n0 | jq -r '.address')
iriscli --home ibc-iris/n0/iriscli q account -o text $(iriscli --home ibc-iris/n0/iriscli keys show n0 | jq -r '.address')