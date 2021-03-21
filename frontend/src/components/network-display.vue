<template>
  <input v-model="networkID" />
</template>
<script>
export default {
  props: ["networkID"],
  emits: ["update:networkID"],
  data() {
    return {
      networkID: "",
    };
  },
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
</script>
