using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class TransactionConfiguration: IEntityTypeConfiguration<Transaction>
{
    public void Configure(EntityTypeBuilder<Transaction> modelBuilder)
    {
        modelBuilder
            .HasKey(t => t.Id)
            .HasName("Tr_id");
    }
}