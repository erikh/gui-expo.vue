import { createApp } from "vue";
import NetworkList from "./components/network-list.vue";
import NetworkDisplay from "./components/network-display.vue";

const RootComponent = {
  components: {
    "network-list": NetworkList,
    "network-display": NetworkDisplay,
  },
  data() {
    return {
      counter: 0,
      networkID: "",
      queriedNetwork: {},
    };
  },
  mounted() {
    setInterval(() => {
      this.counter++;
    }, 1000);
  },
  watch: {
    networkID(newNetwork, old) {
      if (newNetwork.length == 16) {
        getNetwork(newNetwork)
          .then(JSON.parse)
          .then((r) => (this.queriedNetwork = r));
      }
    },
  },
};

createApp(RootComponent).mount("#app");
