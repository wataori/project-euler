import java.util.HashMap;

class Problem017 {
  static HashMap<Integer,String> map = new HashMap<Integer,String>() {
    {
      put(1, "one");
      put(2, "two");
      put(3, "three");
      put(4, "four");
      put(5, "five");
      put(6, "six");
      put(7, "seven");
      put(8, "eight");
      put(9, "nine");
      put(10, "ten");
      put(11, "eleven");
      put(12, "twelve");
      put(13, "thirteen");
      put(14, "fourteen");
      put(15, "fifteen");
      put(16, "sixteen");
      put(17, "seventeen");
      put(18, "eighteen");
      put(19, "nineteen");
      put(20, "twenty");
      put(30, "thirty");
      put(40, "forty");
      put(50, "fifty");
      put(60, "sixty");
      put(70, "seventy");
      put(80, "eighty");
      put(90, "ninety");
      put(1000, "one thousand");
    }
  };

  public static void main(String[] args) {
    int length = 0;
    String number;

    for (int i=1; i<=1000; i++) {
      if (map.containsKey(i)) {
        number = map.get(i);
      } else if (i<100) {
        int second = i/10*10;
        number = map.get(second)+"-"+map.get(i%10);
      } else if (i%100==0) {
        number = map.get(i/100)+" hundred";
      } else if (i<=999) {
        int third = i/100;
        if (map.containsKey(i%100)) {
          number = map.get(third)+" hundred and "+map.get(i%100);
        } else {
          int second = i%100-i%10;
          number = map.get(third)+" hundred and "+map.get(second)+"-"+map.get(i%10);
        }
      } else {
        number = map.get(i);
      }
      length += number.replaceAll("[ -]", "").length();
    }
    System.out.println(length);
  }
}
