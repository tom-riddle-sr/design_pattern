interface Drink{
String printStr();
}

class Coffee implements Drink{
 private int temperature = 0;
 private boolean hasIce = false;
 
 public void setTemperature(int temp) {
   this.temperature = temp;
 }
 
 public void addIce() {
   this.hasIce = true;
 }
 
 public String printStr(){
   return "This is Coffee (temp: " + temperature + ", ice: " + hasIce + ")";
 }
}

class Tea implements Drink{
 private int temperature = 0;
 private boolean hasIce = false;
 
 public void setTemperature(int temp) {
   this.temperature = temp;
 }
 
 public void addIce() {
   this.hasIce = true;
 }
 
 public String printStr(){
   return "This is Tea (temp: " + temperature + ", ice: " + hasIce + ")";
 }
}



class DrinkFactory{
    public static Drink getDrink(String drinkType ){
        if (drinkType.equals("Coffee")) {
            Coffee coffee = new Coffee();
            coffee.setTemperature(85);
            coffee.addIce();
            return coffee;
        }
        else if (drinkType.equals("Tea")){
            Tea tea = new Tea();
            tea.setTemperature(75);
            tea.addIce();
            return tea;
        }
        return null;
    }
}


public class main {
 public static void main(String[] args) {
   Drink drink1 = DrinkFactory.getDrink("Coffee");
   Drink drink2 = DrinkFactory.getDrink("Tea");
   System.out.println(drink1.printStr());
   System.out.println(drink2.printStr());
 }
}