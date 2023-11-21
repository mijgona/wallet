using DataAccess;
using DataAccess.Repositories;

namespace BusinessLogic;

public sealed class UserService : IUserService
{
    private readonly IUserRepository _userRepository;

    public UserService(IUserRepository userRepository)
    {
        _userRepository = userRepository;
    }

    public async Task<User> CreateUser(UserInfo userInfo)
    {
        User newUser = new()
        {
            Name = userInfo.FirstName,
            LastName = userInfo.LastName,
            MiddleName = userInfo.MiddleName,
            Age = userInfo.Age,
            PhoneNumber = userInfo.Phone,
            PassportNumber = userInfo.PassportNumber,
            Gender = userInfo.Gender.ToGenderEnum(),
            
        };

        return await _userRepository.CreateAsync(newUser);;
    }

    public bool RemoveUser(string firstName, string lastName)
    {
        // if (firstName is not { Length: > 0 })
        //     throw new ArgumentNullException(nameof(firstName));
        //
        // if (lastName is not { Length: > 0 })
        //     throw new ArgumentNullException(nameof(lastName));
        //
        // int userIndex = _userRepository.FindIndex(c => c.Name.Equals(firstName) && c.LastName.Equals(lastName));
        // if (userIndex < 0)
        //     return false;
        //
        // _users.RemoveAt(userIndex);

        return true;
    }
    
    //
    // public UserInfo GetUser(string firstName, string lastName)
    // {
    //     // if (firstName is not { Length: > 0 })
    //     //     throw new ArgumentNullException(nameof(firstName));
    //     //
    //     // if (lastName is not { Length: > 0 })
    //     //     throw new ArgumentNullException(nameof(lastName));
    //     //
    //     // User? user = _users.Find(c => c.Name.Equals(firstName) && c.LastName.Equals(lastName))
    //     //     ?? throw new NotFoundException();
    //     return _userRepository.ToUserInfo();
    // }
}
