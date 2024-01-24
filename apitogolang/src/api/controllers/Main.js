// import { connection, queryDatabase } from '../models/UserModel.js';
import { Client } from '@notionhq/client';


class Main {
    init(req, res) {
        const notion = new Client({ auth: NOTION_SECRET });

        (async () => {
            const databaseId = DB_ID;
            const response = await notion.databases.query({
                database_id: databaseId,
                filter: {
                    "property": "Time",
                    "date": {
                        "equals": "2024-01-23"
                    }
                },
                sorts: [
                    {
                        property: 'Time',
                        direction: 'ascending',
                    },
                ],
            });

            const export_data = response.results.map(el => {
                const { IsProductive, Time, Title } = el.properties
                const formattedTime = (lx, z) => {
                    let timeString = lx.date[z].split("T")[1].split('.')[0]
                    const [hours, minutes, seconds] = timeString.split(":");
                    return `${hours}h${minutes}m${seconds}s`;
                }
                return {
                    IsProductive: IsProductive.checkbox,
                    Start: formattedTime(Time, "start"),
                    Title: Title.title[0].plain_text
                }
            })
            console.log(export_data);

            res
                .status(200)
                .send(export_data);
        })();
    }
};

export default new Main();