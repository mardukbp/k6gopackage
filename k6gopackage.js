import k6gopackage from "k6/k6gopackage";

export default function() {
    // On the JS side the method name is in lowercase!
    let result = k6gopackage.method("Hello World!");
    console.log(result);
}