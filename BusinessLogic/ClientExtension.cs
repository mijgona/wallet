using DataAccess;

namespace BusinessLogic;

public static class UserExtension
{
    public static UserInfo ToUserInfo(this User client)
    {
        return new(
            client.Id,
            client.Name,
            client.LastName,
            client.MiddleName,
            client.PhoneNumber,
            client.PassportNumber,
            client.Age,
            client.Gender.ToString());
    }
}
