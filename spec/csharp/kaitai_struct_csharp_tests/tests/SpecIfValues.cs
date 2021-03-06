using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecIfValues : CommonSpec
    {
        [Test]
        public void TestIfValues()
        {
            var r = IfValues.FromFile(SourceFile("fixed_struct.bin"));
            Assert.AreEqual(r.Codes[0].Opcode, 80);
            Assert.AreEqual(r.Codes[0].HalfOpcode, 40);
            Assert.AreEqual(r.Codes[1].Opcode, 65);
            Assert.IsNull(r.Codes[1].HalfOpcode);
            Assert.AreEqual(r.Codes[2].Opcode, 67);
            Assert.IsNull(r.Codes[2].HalfOpcode);
        }
    }
}
