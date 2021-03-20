import { createApp } from "vue";

const RootComponent = {
  data() {
    return {
      counter: 0,
    };
  },
  mounted() {
    setInterval(() => {
      this.counter++;
    }, 1000);
  },
};

createApp(RootComponent).mount("#app");
