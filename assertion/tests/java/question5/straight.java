class Mumble {
    static String mumble(String input) {
        var inputLen = input.length();

        var triNum = (inputLen * (inputLen + 1)) / 2;
        var resultLen = triNum + inputLen - 1;

        var sb = new StringBuilder();
        sb.setLength(resultLen);

        var pos = 0;
        for (int i = 0; i < inputLen; i++) {
            for (int j = 0; j <= i; j++) {
                var c = input.charAt(i);

                var appended = j == 0 ? Character.toUpperCase(c) : Character.toLowerCase(c);
                sb.setCharAt(pos, appended);
                pos++;
            }
            if (pos != resultLen) {
                sb.setCharAt(pos, '-');
                pos++;
            }
        }
        return sb.toString();
    }
}