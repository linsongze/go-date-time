

# Usage
-------------
```
location, _ := time.LoadLocation("UTC")
testTime := Create{time.Date(2024, 6, 3, 10, 15, 30, 0, location)}
formattedTime := testTime.Strftime("%Y-%m-%d %H:%M:%S")
fmt.Println(formattedTime)  // print: 2024-06-03 10:15:30
```

-----------

# Formatting Directives

The format string can include the following directives, each starting with a `%` character:

| Directive | Description |
|-----------|-------------|
| `%a`      | Abbreviated weekday name (e.g., Sun) |
| `%A`      | Full weekday name (e.g., Sunday) |
| `%w`      | Day of the week as a number (0 represents Sunday, 6 represents Saturday) |
| `%d`      | Day of the month as a two-digit number (01 to 31) |
| `%b`      | Abbreviated month name (e.g., Jan) |
| `%B`      | Full month name (e.g., January) |
| `%m`      | Month as a two-digit number (01 to 12) |
| `%y`      | Last two digits of the year |
| `%Y`      | Full year |
| `%H`      | Hour in 24-hour format, as a two-digit number (00 to 23) |
| `%I`      | Hour in 12-hour format, as a two-digit number (01 to 12) |
| `%p`      | AM or PM |
| `%M`      | Minute as a two-digit number (00 to 59) |
| `%S`      | Second as a two-digit number (00 to 59) |
| `%f`      | Microsecond as a six-digit number |
| `%z`      | Timezone as +HHMM or -HHMM (e.g., +0800) |
| `%Z`      | Timezone name (e.g., CST) |
| `%j`      | Day of the year as a three-digit number (001 to 366) |
| `%U`      | Week number of the year, with Sunday as the first day of the week, as a two-digit number (00 to 53) |
| `%W`      | Week number of the year, with Monday as the first day of the week, as a two-digit number (00 to 53) |
| `%c`      | Localized date and time representation, e.g., Mon Jan 2 15:04:05 2006 |
| `%x`      | Localized date representation, e.g., 01/02/06 |
| `%X`      | Localized time representation, e.g., 15:04:05 |
| `%%`      | Percent sign literal |