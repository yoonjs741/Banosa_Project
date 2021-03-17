// 탈 jquery 하잔다.
// jquery가 javacscipt를 편하게 쓰려고 만든 라이브러리인데 
// 요즘 javascipt가 잘 나와서 잘 안쓰는 추세라고 함.
//  Original -> javascript + jquery
// function delHostzone(){
//     var cnt = $("input[name='hostchk']:checked").length;
//     var arr = new Array();
//     $("input[name='hostchk']:checked").each(function() {
//         arr.push($(this).attr('value'));
//     })
//     if(cnt == 0){
//         swal('There is no selected domain')
//     }

// Fixed -> only javascript
// function delHostzone(){
//     var checkedboxes = document.querySelectorAll("input[name='hostchk']:checked");
//     var arr = new Array();
//     for(i=0; i<checkedboxes.length; i++) {
//         arr.push(checkedboxes[i].getAttribute('value'));
//     }
//     if(checkedboxes.length == 0){
//         swal('There is no selected domain')
//     }