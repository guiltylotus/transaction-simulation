abigen --abi ./abi/OneInchRouter.abi --pkg oneinch --type OneInchRouter --out one_inch_router.go &
abigen --abi ./abi/SimSwap.abi --pkg simswap --type SimSwap --out simswap.go  &
mv ./one_inch_router.go contract/one_inch_router &
mv ./simswap.go contract/simswap
