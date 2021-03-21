import { createApp } from "vue";

const NetworkList = {
  data() {
    return {
      list: [],
    };
  },
  template: `
  <ul v-for="n in list">
    <li>{{ n.name }}: {{ n.id }} | {{ n.mac }} | {{ n.portDeviceName }}</li>
  </ul>
  `,
  mounted() {
    setInterval(() => {
      getNetworks()
        .then(JSON.parse)
        .then((r) => (this.list = r));
    }, 1000);
  },
};

const NetworkDisplayComponent = {
  props: ["networkID"],
  emits: ["update:networkID"],
  data() {
    return {
      networkID: "",
    };
  },
  template: `
  <input v-model="networkID" />
  `,
  computed: {
    networkID: {
      get() {
        return this.networkID;
      },
      set(newNetwork) {
        this.$emit("update:networkID", newNetwork);
      },
    },
  },
};

const RootComponent = {
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

app = createApp(RootComponent);
app.component("network-display", NetworkDisplayComponent);
app.component("network-list", NetworkList);
app.mount("#app");
