--  вывести дату в другом формате:
SELECT to_char( current_date, 'dd-mm-yyyy' );

SELECT current_time;

-- извлечь часть даты из строки
SELECT extract( 'year' FROM timestamp '1999-11-27 12:34:56.123459' );