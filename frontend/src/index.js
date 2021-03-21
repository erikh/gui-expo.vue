import { createApp } from "vue";

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
      networkName: "",
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
          .then((r) => JSON.parse(r))
          .then((r) => (this.networkName = r.name));
      }
    },
  },
};

app = createApp(RootComponent);
app.component("network-display", NetworkDisplayComponent);
app.mount("#app");
