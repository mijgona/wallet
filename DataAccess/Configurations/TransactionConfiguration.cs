using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

namespace DataAccess;

public sealed class TransactionConfiguration: IEntityTypeConfiguration<Transaction>
{
    public void Configure(EntityTypeBuilder<Transaction> modelBuilder)
    {
        modelBuilder
            .HasKey(t => t.Id)
            .HasName("pk_id");

        modelBuilder
            .Property(p => p.Id)
            .HasColumnType("SERIAL")
            .HasColumnName("id")
            .IsRequired();

        modelBuilder
            .HasOne(p => p.SourceWallet)
            .WithMany(w => w.Transactions);
        
        modelBuilder
            .HasOne(p => p.TargetWallet)
            .WithMany(w => w.Transactions);
        
        
    }
}