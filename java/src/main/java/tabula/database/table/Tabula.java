package tabula.database.table;

import java.io.BufferedWriter;
import java.io.FileWriter;
import java.io.File;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.text.SimpleDateFormat;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Date;
import java.util.List;
import java.util.Map;

import lombok.AllArgsConstructor;
import lombok.Getter;
import tabula.database.query.QueryUtils;
import tabula.logger.LoggerUtil;

@AllArgsConstructor
@Getter
public class Tabula {
    private String destFolder;
    private int borderStyle;
    private String headerStyleLink;
    private Map<Integer, Header> headers;
    private List<List<String>> rows;

    public void generate() {
        var b = Border.getBorder(this.borderStyle);

        var headerUp = new StringBuilder(b.cornerUpLeft());
        var headerMid = new StringBuilder(b.vertical());
        var headerBottom = new StringBuilder(b.verticalLeft());

        int headersLength = headers.size();
        for (int key = 1; key <= headersLength; key++) {
            int length = headers.get(key).getLength();
            headerUp.append(String.join("", Collections.nCopies(length, b.horizontal())));
            headerBottom.append(String.join("", Collections.nCopies(length, b.horizontal())));
            headerMid.append(addSpaces(headers.get(key).getName(), length)).append(b.vertical());

            if (key < headersLength) {
                headerUp.append(b.divisionUp());
                headerBottom.append(b.intersection());
            } else {
                headerUp.append(b.cornerUpRight());
                headerBottom.append(b.verticalRight());
            }
        }

        List<String> table = new ArrayList<>();
        table.add(headerUp.toString() + "\n");
        table.add(headerMid.toString() + "\n");
        table.add(headerBottom.toString() + "\n");

        var rowsLength = rows.size() - 1;
        var rowFieldsLength = rows.get(0).size() - 1;
        for (int i = 0; i < rows.size(); i++) {
            var row = rows.get(i);
            var value = new StringBuilder(b.vertical());
            var line = new StringBuilder();

            if (i < rowsLength) {
                line.append(b.verticalLeft());
            } else {
                line.append(b.cornerBottomLeft());
            }

            for (int j = 0; j < row.size(); j++) {
                var field = row.get(j);
                value.append(addSpaces(field, headers.get(j + 1).getLength())).append(b.vertical());

                line.append(String.join("", Collections.nCopies(headers.get(j + 1).getLength(), b.horizontal())));
                if (i < rowsLength) {
                    if (j < rowFieldsLength) {
                        line.append(b.intersection());
                    } else {
                        line.append(b.verticalRight());
                    }
                } else if (j < rowFieldsLength) {
                    line.append(b.divisionBottom());
                } else {
                    line.append(b.cornerBottomRight());
                }
            }
            table.add(value.toString() + "\n");
            table.add(line.toString() + "\n");
        }

        var filePath = createTabulaFileFormat(destFolder);
        LoggerUtil.debugf("File path: %s", filePath);
        System.out.println(highlighting(headers, headerStyleLink));
        System.out.println(filePath);

        writeToFile(filePath, table.toArray(new String[0]));
    }

    private String highlighting(Map<Integer, Header> headers, String style) {
        var result = new StringBuilder();
        for (Map.Entry<Integer, Header> entry : headers.entrySet()) {
            var k = entry.getKey();
            var v = entry.getValue();
            result.append(String.format("syn match header%d '%s' | hi link header%d %s |", k, v.getName(), k, style));
        }
        LoggerUtil.debugf("Highlight match: %s", result.toString());
        return result.toString();
    }

    private String addSpaces(String inputString, int length) {
        var result = new StringBuilder(inputString);
        var lengthInputString = QueryUtils.unicodeLength(inputString);

        if (length > lengthInputString) {
            var diff = length - lengthInputString;
            result.append(" ".repeat(diff));
        }

        return result.toString();
    }

    public static void writeToFile(String filePath, String... values) {
        try (BufferedWriter writer = new BufferedWriter(new FileWriter(new File(filePath), StandardCharsets.UTF_8))) {
            for (String v : values) {
                writer.write(v);
            }
        } catch (IOException e) {
            System.out.println("[ERROR] %s".formatted(e.getMessage()));
        }
    }

    public static String createTabulaFileFormat(String destFolder) {
        var sdf = new SimpleDateFormat("yyyyMMdd-HHmmss");
        return String.format("%s/%s.%s", destFolder, sdf.format(new Date()), "tabula");
    }

    public static String createTabulaMongoFileFormat(String destFolder) {
        var sdf = new SimpleDateFormat("yyyyMMdd-HHmmss");
        return String.format("%s/%s.%s.%s", destFolder, sdf.format(new Date()), "tabula", "json");
    }
}
