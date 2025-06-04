async function main() {
  const Recarga = await ethers.getContractFactory("RecargaVE");
  const recarga = await Recarga.deploy();
  await recarga.deployed();
  console.log("Contrato implantado em:", recarga.address);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
