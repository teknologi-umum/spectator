class Grade {
    public static ArrayList<Integer> calculateGrade(ArrayList<Integer> input) {
        var result = new ArrayList<Integer>();
        for (int i = 0; i < input.size(); i++) {
            int grade = input.get(i).intValue();
            if (grade >= 38 && grade % 5 != 0 && grade % 5 >= 3) {
                int added = grade + (5 - grade % 5);
                result.add(added);
            } else {
                result.add(grade);
            }
        }
        return result;
    }
}