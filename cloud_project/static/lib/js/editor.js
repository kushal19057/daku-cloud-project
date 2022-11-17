// Retrieve Elements
const executeCodeBtn = document.querySelector(".editor__run");
const resetCodeBtn = document.querySelector(".editor__reset");

// setup Ace
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

// Events

executeCodeBtn.addEventListener('click', () => {
  // get input from code editor
  const userCode = codeEditor.getValue();

  // run user code
  console.log(userCode);
});

resetCodeBtn.addEventListener('click', () => {
  // clear ace editor
  codeEditor.setValue('');
});

editorLib.init();