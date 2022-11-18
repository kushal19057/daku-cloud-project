let codeEditor = ace.edit("editorCode");

let editorLib = {
    init() {
        // theme
        codeEditor.setTheme("ace/theme/eclipse");

        // set options
        codeEditor.setOptions({
            fontSize: '24pt',
        });
    }   
}


editorLib.init();
        // Retrieve Elements
        const executeCodeBtn = document.querySelector(".editor__run");
        const resetCodeBtn = document.querySelector(".editor__reset");
        const consoleLogs = document.querySelector(".editor__console")

        // Events
        executeCodeBtn.addEventListener('click', () => {
            // get input from code editor
            const userCode = codeEditor.getValue();
            
            // run user code
            console.log(userCode);

            consoleLogs.innerHTML += userCode
        });

        resetCodeBtn.addEventListener('click', () => {
            // clear ace editor
            codeEditor.setValue('');
        });