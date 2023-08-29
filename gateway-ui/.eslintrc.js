module.exports = {
    parserOptions: {
        parser: 'babel-eslint',
        sourceType: 'module'
    },
    plugins: [
        'vue'
    ],
    extends: [
        'plugin:vue/vue3-essential'
    ],
    rules: {
        'vue/multi-word-component-names': 'off'
    }
};
