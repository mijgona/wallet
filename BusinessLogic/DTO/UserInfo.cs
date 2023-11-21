namespace BusinessLogic;

public readonly record struct UserInfo(
    long Id,
    string FirstName,
    string LastName,
    string? MiddleName,
    string Phone,
    string? PassportNumber,
    short Age,
    string Gender);