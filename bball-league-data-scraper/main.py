import requests
import bs4

team_abbreviations = {
            'Brooklyn Nets': 'BRK',     # Switched from NJN to BRK
            'Golden State Warriors': 'GSW',
            'Los Angeles Lakers': 'LAL',
            'Los Angeles Clippers': 'LAC',
            'New Orleans Pelicans': 'NOP',
            'New York Knicks': 'NYK',
            'Oklahoma City Thunder': 'OKC',
            'San Antonio Spurs': 'SAS',
            'Boston Celtics': 'BOS',
            'Denver Nuggets': 'DEN',
            'Minnesota Timberwolves': 'MIN',
            'Cleveland Cavaliers': 'CLE',
            'Philadelphia 76ers': 'PHI',
            'Phoenix Suns': 'PHO',
            'Sacramento Kings': 'SAC',
            'Indiana Pacers': 'IND',
            'Dallas Mavericks': 'DAL',
            'Miami Heat': 'MIA',
            'Orlando Magic': 'ORL',
            'Chicago Bulls': 'CHI',
            'Atlanta Hawks': 'ATL',
            'Toronto Raptors': 'TOR',
            'Charlotte Hornets': 'CHO',     # Switched from CHA to CHO
            'Washington Wizards': 'WAS',
            'Detroit Pistons': 'DET',
            'Utah Jazz': 'UTA',
            'Houston Rockets': 'HOU',
            'Memphis Grizzlies': 'MEM',
            'Portland Trail Blazers': 'POR',
            'Milwaukee Bucks': 'MIL',
}

# for team, abbreviation in team_abbreviations.items():
    # response = requests.get('https://www.basketball-reference.com/teams/' + abbreviation + '/2025.html#all_roster')
response = requests.get('https://www.basketball-reference.com/teams/BOS/2025.html#all_roster')
response.raise_for_status()
soup = bs4.BeautifulSoup(response.text, 'html.parser')

table = soup.find('table', {'id': 'roster'})

trs = table.find_all('tr')

for tr in trs:
    tds = tr.find_all('td')
    if len(tds) > 0:
        # Split into first name and last name
        name_parts = tds[0].find('a').get_text().split(' ')
        first_name = name_parts[0]
        if len(name_parts) > 2:
            last_name = ' '.join(name_parts[1:])
        else:
            last_name = name_parts[1]

        # Because there are no precise position information, I have to scrape positions manually from player's page
        a_tag = tds[0].find('a')
        player_page = 'https://www.basketball-reference.com' + a_tag.get('href')
        player_page_response = requests.get(player_page)
        player_page_response.raise_for_status()
        player_page_soup = bs4.BeautifulSoup(player_page_response.text, 'html.parser')

        position_div = player_page_soup.find('div', {'id': 'meta'})
        p_tags = position_div.find_all('p')
        for p in p_tags:
            if p.get_text().__contains__('Position'):
                if p.get_text().__contains__('Point Guard'):
                    print('Point Guard')
                    # TODO handle PG actions:

                elif p.get_text().__contains__('Shooting Guard'):
                    print('Shooting Guard')
                    # TODO handle SG actions:

                elif p.get_text().__contains__('Small Forward'):
                    print('Small Forward')
                    # TODO handle SF actions:

                elif p.get_text().__contains__('Power Forward'):
                    print('Power Forward')
                    # TODO handle PF actions:

                elif p.get_text().__contains__('Center'):
                    print('Center')
                    # TODO handle C actions:

        # Ovde moram da napravim poseban switch recimo (ili nesto mnogo slicno) pa da tako onda odredim string za to
        position = tds[1].get_text() # Ovaj ne daje dovoljno informacija o poziciji...

        # Convert height from feet and inches to centimeters
        height_parts = tds[2].get_text().split('-')     # Split the string into feet and inches for easier conversion
        height = int(int(height_parts[0]) * 30.48 + int(height_parts[1]) * 2.54)

        # Convert weight from pounds to kilograms
        if len(tds[3].get_text()) > 0:
            weight = str(int(float(tds[3].get_text()) * 0.453592))
        else:
            weight = 'Unknown'  # Means that there is no information about the weight on the basketball-reference page

        # Ovde moram da pretvorim u neki format za datume mislim da ce mi to biti jasno tek kada krenem sa implementacijom bekenda
        birthday = tds[4].get_text()

        # These attributes don't value to scraping that much, so they'll be very simple
        password = last_name + '123'
        email = first_name + '@gmail.com'
