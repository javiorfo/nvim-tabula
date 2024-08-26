package tabula.database.query;

import java.util.List;
import java.util.stream.Stream;

public class QueryUtils {

    public static boolean isSelectQuery(String query) {
        query = query.trim();
        return query.length() == 0 ? false : query.toUpperCase().startsWith("SELECT");
    }

    public static boolean isInsertUpdateOrDelete(String query) {
        query = query.trim();

        if (query.length() == 0) {
            return false;
        }

        if (query.toUpperCase().startsWith("INSERT") || query.toUpperCase().startsWith("UPDATE")
                || query.toUpperCase().startsWith("DELETE")) {
            return true;
        }
        return false;
    }

    public static List<String> splitQueries(String queries) {
        return Stream.of(queries.split(";")).filter(q -> {
            var query = q.trim();
            return !query.isEmpty();
        }).toList();
    }

    public static boolean containsSemicolonInMiddle(String s) {
        var index = s.trim().indexOf(";");
        return index != -1 ? (index > 0 && index < s.length() - 1) : false;
    }

    public static int unicodeLength(String str) {
        return str.codePointCount(0, str.length());
    }
}
