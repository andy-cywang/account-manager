db.createCollection("merchant")

db.merchant.insert(
    {
        _id:"aa584727-ec8c-45c4-a17d-f28ebf509ff9",
        logo:"image1",
        members: {
            jollibee: {
                email: "jollibee@pace.com",
                name: "Jollibee"
            },
            breadtalk: {
                email: "breadtalk@pace.com",
                name: "BreadTalk"
            }
        }
    },
    {
        _id:"fbe206dd-4154-48dd-8d13-6fd065290ca8",
        logo:"image2",
        members: {
            tom: {
                email: "tom@pace.com",
                name: "Tom"
            },
            lisa: {
                email: "lisa@pace.com",
                name: "Lisa"
            },
            jack: {
                email: "jack@pace.com",
                name: "Jack"
            }
        }
    }
)