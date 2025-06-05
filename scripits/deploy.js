const { ethers } = require("hardhat");

async function main() {
  const Recarga = await ethers.getContractFactory("RecargaVE");
  const recarga = await Recarga.deploy();
  await recarga.waitForDeployment();
  console.log("Contrato implantado em:", await recarga.getAddress());
  console.log("Hash da transação de implantação:", recarga.deploymentTransaction().hash);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});