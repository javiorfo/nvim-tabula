package tabula.database.table;

import java.util.HashMap;
import java.util.Map;

public record Border(String cornerUpLeft, String cornerUpRight, String cornerBottomLeft, String cornerBottomRight,
        String divisionUp, String divisionBottom, String horizontal,
        String vertical, String intersection, String verticalLeft, String verticalRight) {

    public enum BorderOption {
        DEFAULT(1),
        SIMPLE(2),
        ROUNDED(3),
        DOUBLE(4),
        SIMPLE_DOUBLE(5);

        private final int value;

        BorderOption(int value) {
            this.value = value;
        }

        public int getValue() {
            return value;
        }

        public static BorderOption get(int value) {
            for (BorderOption bo : BorderOption.values()) {
                if (bo.value == value) {
                    return bo;
                }
            }
            return null;
        }
    }

    private static final Map<BorderOption, Border> borders = new HashMap<>();

    static {
        borders.put(BorderOption.DEFAULT, new Border("┏", "┓", "┗", "┛", "┳", "┻", "━", "┃", "╋", "┣", "┫"));
        borders.put(BorderOption.SIMPLE, new Border("┌", "┐", "└", "┘", "┬", "┴", "─", "│", "┼", "├", "┤"));
        borders.put(BorderOption.ROUNDED, new Border("╭", "╮", "╰", "╯", "┬", "┴", "─", "│", "┼", "├", "┤"));
        borders.put(BorderOption.DOUBLE, new Border("╔", "╗", "╚", "╝", "╦", "╩", "═", "║", "╬", "╠", "╣"));
        borders.put(BorderOption.SIMPLE_DOUBLE, new Border("╒", "╕", "╘", "╛", "╤", "╧", "═", "│", "╪", "╞", "╡"));
    }

    public static Border getBorder(int option) {
        return borders.get(BorderOption.get(option));
    }
}
