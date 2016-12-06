using NUnit.Framework;

namespace Kaitai
{
    [TestFixture]
    public class SpecNavParent2 : CommonSpec
    {
        [Test]
        public void TestNavParent2()
        {
            var r = NavParent2.FromFile(SourceFile("nav_parent2.bin"));
            Assert.AreEqual(r.OfsTags, 8);
            Assert.AreEqual(r.NumTags, 2);
    
            Assert.AreEqual(r.Tags[0].Name, "RAHC");
            Assert.AreEqual(r.Tags[0].Ofs, 0x20);
            Assert.AreEqual(r.Tags[0].NumItems, 3);
            Assert.AreEqual(r.Tags[0].TagContent.Content, "foo");
    
            Assert.AreEqual(r.Tags[1].Name, "RAHC");
            Assert.AreEqual(r.Tags[1].Ofs, 0x23);
            Assert.AreEqual(r.Tags[1].NumItems, 6);
            Assert.AreEqual(r.Tags[1].TagContent.Content, "barbaz");
        }
    }
}
