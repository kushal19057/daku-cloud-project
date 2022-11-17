// Retrieve Elements
const executeCodeBtn = document.querySelector(".editor__run");
const resetCodeBtn = document.querySelector(".editor__reset");

// setup Ace
let codeEditor = ace.edit("editorCode");

let editorLang = document.querySelector('#editorLang');
let defaultMessage = 'Write your code here';

function getLang() {
    return editorLang.value;
}

let editorLib = {
    init() {
        // theme
        codeEditor.setTheme("ace/theme/eclipse");

        // DO NOT SET ANY CODE FAMILY PEHLE SE

        // set options
        codeEditor.setOptions({
            fontSize: '14pt',
            enableBasicAutocompletion: true,
            enableLiveAutocompletion: true,
        });

        // set default message
        codeEditor.setValue(defaultMessage);
    }   
}

// Events
editorLang.addEventListener("change", () => {
  let whichLang = getLang();
  switch (whichLang) {
    case "Java":
        codeEditor.session.setMode("ace/mode/java");
      break;
    case "Python":
        codeEditor.session.setMode("ace/mode/python");
      break;
    case "JavaScript":
        codeEditor.session.setMode("ace/mode/javascript");
      break;
    case "C++":
        codeEditor.session.setMode("ace/mode/c_cpp");
      break;
  }
});

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