require.config({
    paths: {vs: '/vscode/monaco-editor/min/vs'},
    'vs/nls': {availableLanguages: {'*': 'zh-cn'}},
});
require(['vs/editor/editor.main'], function () {
    let editor = window.monaco.editor.create(document.querySelector("#vscode"), {
        value: data,
        language: 'json',
        theme: 'vs-dark',
        fontSize: '16px',
        lineNumbers: 'on',
        roundedSelection: false,
        scrollBeyondLastLine: false,
        readOnly: false,
        tabSize: 4,
        automaticLayout: true,
        readOnlyMessage: true,
        minimap: {
            enabled: true,
            side: 'right',
            showSlider: 'mouseover',
        }
    });
    window.addEventListener('resize', event => {
        editor.layout();
    })
});
