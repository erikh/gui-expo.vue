import { createApp } from "vue";

const NetworkDisplayComponent = {
  data() {
    return {
      networkID: "",
      networkName: "",
    };
  },
  watch: {
    networkID(newNetwork, old) {
      getNetworkName(newNetwork).then((r) => (this.networkName = r));
    },
  },
};

createApp(NetworkDisplayComponent).mount("#app");
