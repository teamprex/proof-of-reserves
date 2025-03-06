# Flipster Proof-of-Reserves

## Background

At Flipster, we prioritize transparency and security to ensure that your assets are always fully backed and protected.
Proof of Reserves (PoR) is a verification method that ensures Flipster holds 100% of user deposits on a 1:1 basis.
We provide daily data updates and conduct independent audits to reinforce trust and prove that all user funds are securely held.

This tool allows you to verify PoR data is properly include your assets.

## How to run source code

Place the downloaded `merkle_proof.json` into `data/` directory.
And run the command below.

```
go run main.go
```

## How to run binary

Download the binary from Release according to user operation system.

Place the downloaded `merkle_proof.json` into `data/` directory.
And run the command below.

```
./verifier
```
