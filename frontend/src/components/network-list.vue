<template>
  <ul v-for="n in list">
    <li>{{ n.name }}: {{ n.id }} | {{ n.mac }} | {{ n.portDeviceName }}</li>
  </ul>
</template>
<script>
export default {
  data() {
    return {
      list: [],
    };
  },
  mounted() {
    setInterval(() => {
      getNetworks()
        .then(JSON.parse)
        .then((r) => (this.list = r));
    }, 1000);
  },
};
</script>
