package tabula;

import org.junit.Test;
import tabula.database.query.QueryUtils;
import static org.junit.Assert.*;

import java.util.List;

public class QueryUtilsTest {

    @Test
    public void testIsSelect() {
        assertTrue(QueryUtils.isSelectQuery("select * from table;"));
        assertFalse(QueryUtils.isSelectQuery("insert * from table;"));
        assertFalse(QueryUtils.isSelectQuery(""));
    }
    
    @Test
    public void testIsInsertUpdateOrDelete() {
        assertTrue(QueryUtils.isInsertUpdateOrDelete("insert into table;"));
        assertFalse(QueryUtils.isInsertUpdateOrDelete("create insert * from table;"));
        assertFalse(QueryUtils.isInsertUpdateOrDelete(""));
    }

    @Test
    public void testSplitQueries() {
        var result = List.of("select * from table", "select * from table2");
        assertEquals(QueryUtils.splitQueries("select * from table;select * from table2"), result);
        assertEquals(QueryUtils.splitQueries("insert * from table;"), List.of("insert * from table"));
        assertEquals(QueryUtils.splitQueries(""), List.of());
    }
   
    @Test
    public void testContainsSemicolonInMiddle() {
        assertFalse(QueryUtils.containsSemicolonInMiddle("select * from table"));
        assertFalse(QueryUtils.containsSemicolonInMiddle("insert * from table;"));
        assertTrue(QueryUtils.containsSemicolonInMiddle("insert into table; delete from table;"));
    }
}
