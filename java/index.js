let Eq =document.getElementById("equal-el")
 function calcNumbers(result){
 calcForm.displayResult.value=calcForm.displayResult.value+result;
 }
 function Clear(){
   calcForm.displayResult.value=" ";
 }
 function equalTo(){
   calcForm.displayResult.value = eval(calcForm.displayResult.value)
 }
 function Delet(){
   let Del = document.calcForm.displayResult.value;
   document.calcForm.displayResult.value = Del.substring(0, Del.length - 1);
 }

 function startTime() {
  var today = new Date();
  var hr = today.getHours();
  var min = today.getMinutes();
  var sec = today.getSeconds();
  ap = (hr < 12) ? "<span>AM</span>" : "<span>PM</span>";
  hr = (hr == 0) ? 12 : hr;
  hr = (hr > 12) ? hr - 12 : hr;
  //Add a zero in front of numbers<10
  hr = checkTime(hr);
  min = checkTime(min);
  sec = checkTime(sec);
  document.getElementById("clock").innerHTML = hr + ":" + min + ":" + sec + " " + ap;
  
  var months = ['Janvier', 'Fevrier', 'Mars', 'Avril', 'Mai', 'Juin', 'Juillet', 'Août', 'Septembre', 'Octobre', 'Novembre', 'Decembre'];
  var days = ['Dimanche', 'Lundi', 'Mardi', 'Mercredi', 'Jeudi', 'Vendredi', 'Samedi'];
  var curWeekDay = days[today.getDay()];
  var curDay = today.getDate();
  var curMonth = months[today.getMonth()];
  var curYear = today.getFullYear();
  var date = curWeekDay+", "+curDay+" "+curMonth+" "+curYear;
  document.getElementById("date").innerHTML = date;
  
  var time = setTimeout(function(){ startTime() }, 500);
}
function checkTime(i) {
  if (i < 10) {
      i = "0" + i;
  }
  return i;
}
function calculer(event) {
  event.preventDefault();
  const form1 = document.getElementById("form1");
  const valeur_bien = parseFloat(form1.elements["valeur_bien"].value);
  const taux_commission = parseFloat(form1.elements["taux_commission"].value);
  const frais_agence = (valeur_bien * taux_commission) / 100.0;
  const net_vendeur = (valeur_bien - frais_agence)
  const resultat = document.getElementById("resultat");
  resultat.innerHTML = `Prix affiché: ${valeur_bien.toFixed(0)}€<br>Pourcentage: ${taux_commission.toFixed(0)}%<br>Frais d'agence: ${frais_agence.toFixed(0)}€<br>Net vendeur: ${net_vendeur.toFixed(0)}€.`;
}
function calcul(event) {
  event.preventDefault();
  const form2 = document.getElementById("form2");
  const valeur_all = parseFloat(form2.elements["valeur_all"].value);
  const frais_notaire = (valeur_all + ((valeur_all * 8)/100))
  const result = document.getElementById("result");
  result.innerHTML = `Le prix frais de notaire inclus est de: ${frais_notaire.toFixed(0)}€`
}

//